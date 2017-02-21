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
		} else if(this.type == ElementType.POLYGON) {
			return "polygon";
		} else if(this.type == ElementType.OVAL) {
			return "oval";
		} else {
			return "undefined";
		}
	}
	
	public String toString() {
		String ret = "";
		for(Integer a: this.arguments) {
			ret += " "+Integer.toString(a);
		}
		return ret;
	}
	
	public void draw(Graphics2D g2d) {
		g2d.setColor(Color.black);
	}
}
