package autodraw;

import java.awt.Color;
import java.awt.Graphics2D;
import java.awt.Point;

public class LineElement extends Element {
	private LineElement() {
		
	}
	
	public LineElement(int x1, int y1, int x2, int y2) {
		super();
		this.type = ElementType.LINE;
		this.arguments.add(x1);
		this.arguments.add(y1);
		this.arguments.add(x2);
		this.arguments.add(y2);
	}
	
	public void draw(Graphics2D g2d) {
		g2d.setColor(Color.black);
		g2d.drawLine(this.arguments.get(0), this.arguments.get(1),
				this.arguments.get(2), this.arguments.get(3));
	}

	@Override
	public void drawHighlight(Graphics2D g2d) {
		g2d.setColor(Color.red);
		g2d.drawLine(this.arguments.get(0), this.arguments.get(1),
				this.arguments.get(2), this.arguments.get(3));
	}
	
	@Override
	public Object clone() throws CloneNotSupportedException {
		LineElement e = new LineElement();
		e.type = this.type;
		e.arguments.addAll(this.arguments);
		return e;
	}
	
	public String getType() {
		return "line";
	}
	
	public LineElement translated(int originx, int originy)
			throws CloneNotSupportedException {
		LineElement e = (LineElement)this.clone();
		e.translate(originx, originy);
		return e;
	}

	@Override
	public Touch getTouch(int x, int y) {
		int x0 = this.arguments.get(0), y0 = this.arguments.get(1),
			x1 = this.arguments.get(2), y1 = this.arguments.get(3);
		if(Touch.close(x,y,x0,y0)) return new Touch(this,1);
		if(Touch.close(x,y,x1,y1)) return new Touch(this,2);
		return null;
	}

	@Override
	public Point getPointTouch(int index) {
		int x0 = this.arguments.get(0), y0 = this.arguments.get(1),
			x1 = this.arguments.get(2), y1 = this.arguments.get(3);
		switch(index) {
		case 1: return new Point(x0,y0);
		case 2: return new Point(x1,y1);
		}
		return null;
	}

}
