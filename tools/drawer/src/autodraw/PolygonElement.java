package autodraw;

import java.awt.Graphics2D;
import java.awt.Point;
import java.util.ArrayList;

public class PolygonElement extends Element {
	public PolygonElement(ArrayList<Point> args) {
		super();
		this.type = ElementType.POLYGON;
		for(Point p: args) {
			this.arguments.add(p.x);
			this.arguments.add(p.y);
		}
	}
	
	public void draw(Graphics2D g2d) {
		super.draw(g2d);
		int n = this.arguments.size();
		for(int i = 0; i < n/2-1; i++) {
			g2d.drawLine(this.arguments.get(i*2),this.arguments.get(i*2+1),
					this.arguments.get(i*2+2),this.arguments.get(i*2+3));
		}
		g2d.drawLine(this.arguments.get(n-2), this.arguments.get(n-1),
				this.arguments.get(0), this.arguments.get(1));
	}
	
	public String getType() {
		return "polygon";
	}
}
