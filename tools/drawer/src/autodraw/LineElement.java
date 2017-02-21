package autodraw;

import java.awt.Graphics2D;

public class LineElement extends Element {
	public LineElement(int x1, int y1, int x2, int y2) {
		super();
		this.type = ElementType.LINE;
		this.arguments.add(x1);
		this.arguments.add(y1);
		this.arguments.add(x2);
		this.arguments.add(y2);
	}
	
	public void draw(Graphics2D g2d) {
		super.draw(g2d);
		g2d.drawLine(this.arguments.get(0), this.arguments.get(1),
				this.arguments.get(2), this.arguments.get(3));
	}
}
