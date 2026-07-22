import type { LayoutRect } from './types';

// padding ensures visual breathing room between organic placements
export function isColliding(a: LayoutRect, b: LayoutRect, padding: number = 30): boolean {
	return !(
		a.x + a.width + padding < b.x ||
		b.x + b.width + padding < a.x ||
		a.y + a.height + padding < b.y ||
		b.y + b.height + padding < a.y
	);
}
