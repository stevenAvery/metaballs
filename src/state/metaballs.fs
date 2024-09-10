#version 330

#define FLOAT_MAX         3.402823466e+38
#define AA_SPREAD         1.0  // Spread of Anti-Aliased area
#define MAX_NUM_METABALLS 64

// Input from vertex shader
in vec2 fragTexCoord; // Range from 0 to 1 in positive x and positive y
in vec4 fragColour;   // Always vec4(1.0, 1.0, 1.0, 1.0)

// Output fragment colour
out vec4 finalColour;

// Uniforms
uniform vec2  windowSize;
uniform vec3  metaballs[MAX_NUM_METABALLS];
uniform int   numMetaballs = 0;
uniform vec3  backgroundColour = vec3(0.0, 0.0, 0.0);
uniform vec3  metaballColour = vec3(1.0, 1.0, 1.0);
uniform float blobiness = 0.5; // Higher number makes metaballs more "bloby". Recommended 0.5 to 1.0

// Smooth minimum
float smin(float a, float b, float k) {
	float h = max(k-abs(a-b), 0.0)/k;
	return min(a, b) - h*h*k*(1.0/4.0);
}

// Signed distance function of a circle	
float sdfCircle(vec2 pos1, vec2 pos2, float r) {
	return length(pos2-pos1) - r;
}

void main() {
	vec2 windowCoord = fragTexCoord * windowSize;

	// Calculate metaballs
	float dist = FLOAT_MAX;
	for (int i = 0; i < numMetaballs; i++) {
		float circle = sdfCircle(windowCoord, metaballs[i].xy, metaballs[i].z);

		dist = smin(dist, circle, blobiness*metaballs[i].z);
	}

	float alpha = clamp(0.5 - dist / AA_SPREAD, 0.0, 1.0);

	finalColour = vec4(backgroundColour, 1);
	finalColour = mix(finalColour, vec4(metaballColour, 1.0), alpha);
}