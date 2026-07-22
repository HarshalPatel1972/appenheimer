import type { AppResult } from '../types/search';

export class HoverState {
	hoveredApp = $state<AppResult | null>(null);
	anchorRect = $state<DOMRect | null>(null);

	private timeoutId: ReturnType<typeof setTimeout> | null = null;

		setHover(app: AppResult, rect: DOMRect) {
		if (typeof window !== 'undefined' && !window.matchMedia('(hover: hover) and (pointer: fine)').matches) {
			return; // Disable hover entirely on touch devices
		}
		if (this.timeoutId) clearTimeout(this.timeoutId);
		this.timeoutId = setTimeout(() => {
			this.hoveredApp = app;
			this.anchorRect = rect;
		}, 150);
	}

	clearHover(appId: string) {
		if (this.timeoutId) clearTimeout(this.timeoutId);
		this.timeoutId = setTimeout(() => {
			if (this.hoveredApp?.id === appId) {
				this.hoveredApp = null;
				this.anchorRect = null;
			}
		}, 50); // slight debounce for moving cursor
	}
	keepHoverAlive(appId: string) {
		if (this.timeoutId && this.hoveredApp?.id === appId) {
			clearTimeout(this.timeoutId);
		}
	}
}

export const hoverStore = new HoverState();
