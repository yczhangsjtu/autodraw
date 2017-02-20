package autodraw;

import java.util.ArrayList;

public class Element {
	
	public enum ElementType {
		UNDEFINED, LINE, RECT, POLYGON, OVAL
	}

	private ElementType type;
	private ArrayList<Double> arguments;
	
	public Element() {
		this.type = ElementType.UNDEFINED;
		this.arguments = new ArrayList<Double>();
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
		for(Double a: this.arguments) {
			ret += " "+Double.toString(a);
		}
		return ret;
	}
}
