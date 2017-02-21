package autodraw;

import java.awt.Graphics2D;

public class OvalElement extends Element {
	public OvalElement(int x, int y, int a, int b) {
		super();
		this.type = ElementType.LINE;
		this.arguments.add(x);
		this.arguments.add(y);
		this.arguments.add(a);
		this.arguments.add(b);
	}
	
	public void draw(Graphics2D g2d) {
		super.draw(g2d);
		g2d.drawOval(this.arguments.get(0)-this.arguments.get(2),
				this.arguments.get(1)-this.arguments.get(3),
				this.arguments.get(2)*2, this.arguments.get(3)*2);
	}
}
