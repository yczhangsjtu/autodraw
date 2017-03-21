package autodraw;

import java.awt.Color;
import java.awt.Graphics2D;

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
}
