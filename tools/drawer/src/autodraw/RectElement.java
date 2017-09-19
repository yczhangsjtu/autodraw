package autodraw;

import java.awt.Color;
import java.awt.Graphics2D;
import java.awt.Point;

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
	@Override
	public void drawHighlight(Graphics2D g2d) {
		g2d.setColor(Color.red);
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
	@Override
	public Touch getTouch(int x, int y) {
		int x0 = this.arguments.get(0), y0 = this.arguments.get(1),
			x1 = this.arguments.get(2), y1 = this.arguments.get(3);
		if(Touch.close(x,y,x0,y0)) return new Touch(this,1);
		if(Touch.close(x,y,x1,y1)) return new Touch(this,2);
		if(Touch.close(x,y,x0,y1)) return new Touch(this,3);
		if(Touch.close(x,y,x1,y0)) return new Touch(this,4);
		if(x > x0 && x < x1 && y > y0 && y < y1) return new Touch(this);
		return null;
	}
	@Override
	public Point getPointTouch(int index) {
		int x0 = this.arguments.get(0), y0 = this.arguments.get(1),
			x1 = this.arguments.get(2), y1 = this.arguments.get(3);
		switch(index) {
		case 1: return new Point(x0,y0);
		case 2: return new Point(x1,y1);
		case 3: return new Point(x0,y1);
		case 4: return new Point(x1,y0);
		}
		return null;
	}
}
