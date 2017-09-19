package autodraw;

import java.awt.Color;
import java.awt.Graphics2D;
import java.awt.Point;

public class OvalElement extends Element {
	private OvalElement() {
		
	}
	public OvalElement(int x, int y, int a, int b) {
		super();
		this.type = ElementType.OVAL;
		this.arguments.add(x);
		this.arguments.add(y);
		this.arguments.add(Math.abs(a));
		this.arguments.add(Math.abs(b));
	}
	
	public void draw(Graphics2D g2d) {
		g2d.setColor(Color.black);
		g2d.drawOval(this.arguments.get(0)-this.arguments.get(2),
				this.arguments.get(1)-this.arguments.get(3),
				this.arguments.get(2)*2, this.arguments.get(3)*2);
	}
	@Override
	public void drawHighlight(Graphics2D g2d) {
		g2d.setColor(Color.red);
		g2d.drawOval(this.arguments.get(0)-this.arguments.get(2),
				this.arguments.get(1)-this.arguments.get(3),
				this.arguments.get(2)*2, this.arguments.get(3)*2);
	}
	
	public String getType() {
		return "oval";
	}

	
	@Override
	public Object clone() throws CloneNotSupportedException {
		OvalElement e = new OvalElement();
		e.type = this.type;
		e.arguments.addAll(this.arguments);
		return e;
	}
	
	public OvalElement translated(int originx, int originy) {
		OvalElement e = new OvalElement();
		e.type = this.type;
		e.arguments.addAll(this.arguments);
		e.arguments.set(0, e.arguments.get(0)-originx);
		e.arguments.set(1, originy-e.arguments.get(1));
		return e;
	}
	@Override
	public Touch getTouch(int x, int y) {
		int x0 = arguments.get(0), y0 = arguments.get(1), a = arguments.get(2), b = arguments.get(3);
		int x1 = x0 + a, y1 = y0, x2 = x0 - a, y2 = y0, x3 = x0, y3 = y0 + b, x4 = x0, y4 = y0 - b;
		if(Touch.close(x,y,x1,y1)) return new Touch(this,1);
		if(Touch.close(x,y,x2,y2)) return new Touch(this,2);
		if(Touch.close(x,y,x3,y3)) return new Touch(this,3);
		if(Touch.close(x,y,x4,y4)) return new Touch(this,4);
		if((x-x0)*(x-x0)*b*b+(y-y0)*(y-y0)*a*a<a*a*b*b) return new Touch(this);
		return null;
	}
	@Override
	public Point getPointTouch(int index) {
		int x0 = arguments.get(0), y0 = arguments.get(1), a = arguments.get(2), b = arguments.get(3);
		int x1 = x0 + a, y1 = y0, x2 = x0 - a, y2 = y0, x3 = x0, y3 = y0 + b, x4 = x0, y4 = y0 - b;
		switch(index) {
		case 1: return new Point(x1,y1);
		case 2: return new Point(x2,y2);
		case 3: return new Point(x3,y3);
		case 4: return new Point(x4,y4);
		}
		return null;
	}
}
