import { generateSeed, mulberry32 } from './seed';
import { findPlacement } from './placement';
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

	const placed: PlacedResult[] = [];
	const cardWidth = 260;
	const cardHeight = 84;

	for (const result of results) {
		const targetX = (canvasWidth / 2 - cardWidth / 2) + (random() - 0.5) * canvasWidth * 0.4;
		const targetY = (canvasHeight / 2 - cardHeight / 2) + (random() - 0.5) * canvasHeight * 0.4;

		// Map placed items to LayoutRect for collision detection
		const rects = placed.map(p => p.layout);

		const pos = findPlacement(
			{ x: targetX, y: targetY },
			cardWidth,
			cardHeight,
			rects,
			canvasWidth,
			canvasHeight
		);

		placed.push({
			app: result,
			layout: {
				x: pos.x,
				y: pos.y,
				width: cardWidth,
				height: cardHeight
			}
		});
	}

	return placed;
}
