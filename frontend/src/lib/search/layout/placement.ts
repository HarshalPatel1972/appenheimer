import { isColliding } from './collision';
import type { LayoutRect, Point } from './types';

// Spirals outward to find empty deterministic space
export function findPlacement(
	target: Point,
	width: number,
	height: number,
	placed: LayoutRect[],
	canvasWidth: number,
	canvasHeight: number
): Point {
	let angle = 0;
	let radius = 0;
	const step = 25; // How fast the spiral expands

	while (true) {
		const x = target.x + radius * Math.cos(angle);
		const y = target.y + radius * Math.sin(angle);

		// Clamp to canvas to prevent horizontal/vertical scrolling
		const clampedX = Math.max(20, Math.min(x, canvasWidth - width - 20));
		const clampedY = Math.max(100, Math.min(y, canvasHeight - height - 20)); // Keep below search bar

		const rect: LayoutRect = { x: clampedX, y: clampedY, width, height };

		let hasCollision = false;
		for (const p of placed) {
			if (isColliding(rect, p)) {
				hasCollision = true;
				break;
			}
		}

		if (!hasCollision || radius > Math.max(canvasWidth, canvasHeight)) {
			return { x: clampedX, y: clampedY };
		}

		angle += 0.8;
		radius += step;
	}
}
