export interface HoverPanelPosition {
	x: number;
	y: number;
}

export function calculateHoverPosition(
	anchorRect: DOMRect,
	panelWidth: number,
	panelHeight: number,
	canvasWidth: number,
	canvasHeight: number
): HoverPanelPosition {
	const gap = 20;
	let x = anchorRect.right + gap;
	let y = anchorRect.top + (anchorRect.height / 2) - (panelHeight / 2);

	if (x + panelWidth > canvasWidth) {
		x = anchorRect.left - panelWidth - gap;
	}
	
	x = Math.max(gap, Math.min(x, canvasWidth - panelWidth - gap));
	y = Math.max(gap + 80, Math.min(y, canvasHeight - panelHeight - gap));

	return { x, y };
}
