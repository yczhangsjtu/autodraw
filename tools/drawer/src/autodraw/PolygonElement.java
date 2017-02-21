package autodraw;

import java.awt.Graphics2D;
import java.util.ArrayList;

public class PolygonElement extends Element {
	public PolygonElement(ArrayList<Integer> args) {
		super();
		this.type = ElementType.RECT;
		this.arguments.addAll(args);
	}
	
	public void draw(Graphics2D g2d) {
		super.draw(g2d);
	}
}
