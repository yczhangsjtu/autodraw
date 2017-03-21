package autodraw;

import java.awt.Color;
import java.awt.Graphics2D;

public class RectElement extends Element {
	private RectElement() {
		
	}
	public RectElement(int x1, int y1, int x2, int y2) {
		super();
		this.type = ElementType.RECT;
		if(x1 > x2) {
			int tmp = x1; x1 = x2; x2 = tmp;
		}
		if(y1 > y2) {
			int tmp = y1; y1 = y2; y2 = tmp;
		}
		this.arguments.add(x1);
		this.arguments.add(y1);
		this.arguments.add(x2);
		this.arguments.add(y2);
	}
	
	public void draw(Graphics2D g2d) {
		g2d.setColor(Color.black);
		g2d.drawRect(this.arguments.get(0), this.arguments.get(1),
				this.arguments.get(2)-this.arguments.get(0),
				this.arguments.get(3)-this.arguments.get(1));
	}
	
	public String getType() {
		return "rect";
	}

	
	@Override
	public Object clone() throws CloneNotSupportedException {
		RectElement e = new RectElement();
		e.type = this.type;
		e.arguments.addAll(this.arguments);
		return e;
	}
	
	public RectElement translated(int originx, int originy)
			throws CloneNotSupportedException {
		RectElement e = (RectElement)this.clone();
		e.translate(originx, originy);
		return e;
	}
}
