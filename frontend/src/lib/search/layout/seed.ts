export function generateSeed(query: string): number {
	let h = 0xdeadbeef;
	for (let i = 0; i < query.length; i++) {
		h = Math.imul(h ^ query.charCodeAt(i), 2654435761);
	}
	return (h ^ h >>> 16) >>> 0;
}

// Mulberry32 PRNG
export function mulberry32(a: number) {
	return function() {
		var t = a += 0x6D2B79F5;
		t = Math.imul(t ^ t >>> 15, t | 1);
		t ^= t + Math.imul(t ^ t >>> 7, t | 61);
		return ((t ^ t >>> 14) >>> 0) / 4294967296;
	}
}
