import { generateSeed, mulberry32 } from './seed';
import type { AppResult } from '../../types/search';
import type { PlacedResult } from './types';

export function calculateLayout(
	query: string,
	results: AppResult[],
	canvasWidth: number,
	canvasHeight: number
): PlacedResult[] {
	if (!canvasWidth || !canvasHeight || results.length === 0) return [];

	const seed = generateSeed(query);
	const random = mulberry32(seed);

	const cardWidth = 260;
	const cardHeight = 84;
	const padX = 24;
	const padY = 24;
	const cellW = cardWidth + padX;
	const cellH = cardHeight + padY;

	const centerX = canvasWidth / 2 - cardWidth / 2;
	const centerY = canvasHeight / 2 - cardHeight / 2;

	interface GridPos { q: number; r: number; x: number; y: number; dist: number }
	let grid: GridPos[] = [];

	// Generate grid cells
	const rings = Math.ceil(Math.sqrt(results.length)) + 2;
	for (let r = -rings; r <= rings; r++) {
		for (let q = -rings; q <= rings; q++) {
			const x = q * cellW + (Math.abs(r) % 2 === 1 ? cellW / 2 : 0);
			const y = r * cellH;
			const dist = Math.sqrt(x * x + y * y);
			grid.push({ q, r, x, y, dist });
		}
	}

	// Sort by distance from center, then introduce deterministic organic shuffling for equal distances
	grid.sort((a, b) => {
		if (Math.abs(a.dist - b.dist) < 0.1) {
			return random() - 0.5; // Shuffle cells that are roughly equidistant to make it organic
		}
		return a.dist - b.dist;
	});

	// To make the overall shape slightly asymmetric but stable, we can shuffle chunks
	for (let i = 0; i < grid.length; i += 4) {
		const chunk = grid.slice(i, i + 4);
		for (let j = chunk.length - 1; j > 0; j--) {
			const k = Math.floor(random() * (j + 1));
			[chunk[j], chunk[k]] = [chunk[k], chunk[j]];
		}
		for (let j = 0; j < chunk.length; j++) {
			grid[i + j] = chunk[j];
		}
	}

	const placed: PlacedResult[] = [];
	for (let i = 0; i < results.length; i++) {
		const pos = grid[i];
		placed.push({
			app: results[i],
			layout: {
				x: centerX + pos.x,
				y: centerY + pos.y,
				width: cardWidth,
				height: cardHeight
			}
		});
	}

	return placed;
}
