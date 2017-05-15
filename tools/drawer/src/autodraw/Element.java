package autodraw;

import java.awt.Graphics2D;
import java.util.ArrayList;

abstract public class Element {
	
	public enum ElementType {
		UNDEFINED, LINE, RECT, POLYGON, OVAL, TEXT
	}

	protected ElementType type;
	protected ArrayList<Integer> arguments;
	
	public Element() {
		this.type = ElementType.UNDEFINED;
		this.arguments = new ArrayList<Integer>();
	}
	
	public abstract String getType();
	public abstract void draw(Graphics2D g2d);
	public abstract Element translated(int originx, int originy)
			throws CloneNotSupportedException;
	
	public String toString() {
		String ret = getType();
		for(Integer a: this.arguments) {
			ret += " "+Integer.toString(a);
		}
		return ret;
	}
	
	public void translate(int originx, int originy) {
		for(int i = 0; i < this.arguments.size(); i++) {
			if(i%2 == 0) this.arguments.set(i,this.arguments.get(i)-originx);
			else this.arguments.set(i,originy-this.arguments.get(i));
		}
	}

}
