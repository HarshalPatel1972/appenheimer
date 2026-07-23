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

	const cardWidth = 120;
	const cardHeight = 120;
	
	// Increase padding for a wider spread
	const padX = 60;
	const padY = 60;
	const cellW = cardWidth + padX;
	const cellH = cardHeight + padY;

	const centerX = canvasWidth / 2 - cardWidth / 2;
	const centerY = canvasHeight / 2 - cardHeight / 2;

	// Dead zone in the center for the search bar (approx 850x350)
	const exclusionWidth = 850;
	const exclusionHeight = 350;

	interface GridPos { q: number; r: number; x: number; y: number; dist: number }
	let grid: GridPos[] = [];

	// Generate grid cells
	const rings = Math.ceil(Math.sqrt(results.length)) + 4; // Generate more to account for exclusion
	for (let r = -rings; r <= rings; r++) {
		for (let q = -rings; q <= rings; q++) {
			const x = q * cellW + (Math.abs(r) % 2 === 1 ? cellW / 2 : 0);
			const y = r * cellH;
			
			// Check if cell is inside exclusion zone
			// Since x, y are relative to 0,0 center
			const cardLeft = x - cardWidth / 2;
			const cardRight = x + cardWidth / 2;
			const cardTop = y - cardHeight / 2;
			const cardBottom = y + cardHeight / 2;
			
			const exclLeft = -exclusionWidth / 2;
			const exclRight = exclusionWidth / 2;
			const exclTop = -exclusionHeight / 2;
			const exclBottom = exclusionHeight / 2;
			
			const overlapsX = cardRight > exclLeft && cardLeft < exclRight;
			const overlapsY = cardBottom > exclTop && cardTop < exclBottom;
			
			if (overlapsX && overlapsY) {
				continue; // Skip this cell, it hits the search bar
			}

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
		if (i >= grid.length) break; // Should not happen given +4 rings
		const pos = grid[i];
		
		// Add organic jitter to break the rigid grid, max jitter ±35px
		const jitterX = (random() * 70) - 35;
		const jitterY = (random() * 70) - 35;
		
		const targetX = centerX + pos.x + jitterX;
		const targetY = centerY + pos.y + jitterY;
		
		const clampedX = Math.max(20, Math.min(canvasWidth - cardWidth - 20, targetX));
		const clampedY = Math.max(20, Math.min(canvasHeight - cardHeight - 20, targetY));
		
		placed.push({
			app: results[i],
			layout: {
				x: clampedX,
				y: clampedY,
				width: cardWidth,
				height: cardHeight
			}
		});
	}

	return placed;
}
