package autodraw;

import java.awt.Color;
import java.awt.Graphics2D;
import java.util.ArrayList;

public class Element {
	
	public enum ElementType {
		UNDEFINED, LINE, RECT, POLYGON, OVAL
	}

	protected ElementType type;
	protected ArrayList<Integer> arguments;
	
	public Element() {
		this.type = ElementType.UNDEFINED;
		this.arguments = new ArrayList<Integer>();
	}
	
	public String getType() {
		if(this.type == ElementType.LINE) {
			return "line";
		} else if(this.type == ElementType.RECT) {
			return "rect";
		} else if(this.type == ElementType.OVAL) {
			return "oval";
		} else if(this.type == ElementType.POLYGON) {
			return "polygon";
		} else {
			return "undefined";
		}
	}
	
	public String toString() {
		String ret = getType();
		for(Integer a: this.arguments) {
			ret += " "+Integer.toString(a);
		}
		return ret;
	}
	
	public void draw(Graphics2D g2d) {
		g2d.setColor(Color.black);
	}
	
	public Element translate(int originx, int originy) {
		Element e = new Element();
		e.type = this.type;
		for(int i = 0; i < this.arguments.size(); i++) {
			if(i%2 == 0) e.arguments.add(this.arguments.get(i)-originx);
			else e.arguments.add(originy-this.arguments.get(i));
		}
		return e;
	}
}
