import type { AppResult } from '../../types/search';

export interface Point {
	x: number;
	y: number;
}

export interface LayoutRect extends Point {
	width: number;
	height: number;
}

export interface PlacedResult {
	app: AppResult;
	layout: LayoutRect;
}
