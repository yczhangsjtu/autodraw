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
	
	@Override
	public Object clone() throws CloneNotSupportedException {
		Element e = new Element();
		e.type = this.type;
		e.arguments.addAll(this.arguments);
		return e;
	}
	
	public String getType() {
		return "undefined";
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
	
	public void translate(int originx, int originy) {
		for(int i = 0; i < this.arguments.size(); i++) {
			if(i%2 == 0) this.arguments.set(i,this.arguments.get(i)-originx);
			else this.arguments.set(i,originy-this.arguments.get(i));
		}
	}

	public Element translated(int originx, int originy)
			throws CloneNotSupportedException {
		Element e = (Element)this.clone();
		e.translate(originx, originy);
		return e;
	}
}
