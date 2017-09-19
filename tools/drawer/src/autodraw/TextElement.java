package autodraw;

import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics2D;
import java.awt.Point;

public class TextElement extends Element {
	private final String text;
	private int w, h;
	public TextElement(String text, int x, int y) {
		super();
		this.type = ElementType.TEXT;
		this.text = text;
		this.arguments.add(x);
		this.arguments.add(y);
	}
	@Override
	public String getType() {
		return "text";
	}

	@Override
	public void draw(Graphics2D g2d) {
		g2d.setColor(Color.black);
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,24));
		this.w = g2d.getFontMetrics().stringWidth(this.text);
		this.h = g2d.getFontMetrics().getHeight();
		g2d.drawString(this.text, this.arguments.get(0)-w/2, this.arguments.get(1)+h/2);
	}
	@Override
	public void drawHighlight(Graphics2D g2d) {
		g2d.setColor(Color.red);
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,24));
		int w = g2d.getFontMetrics().stringWidth(this.text);
		int h = g2d.getFontMetrics().getHeight();
		g2d.drawString(this.text, this.arguments.get(0)-w/2, this.arguments.get(1)+h/2);
	}

	@Override
	public TextElement translated(int originx, int originy)
			throws CloneNotSupportedException {
		TextElement e = (TextElement)this.clone();
		e.translate(originx, originy);
		return e;
	}
	
	@Override
	public Object clone() throws CloneNotSupportedException {
		TextElement e = new TextElement(this.text,this.arguments.get(0),this.arguments.get(1));
		return e;
	}
	
	public String toString() {
		return String.format("text %d %d %d \"%s\"", this.arguments.get(0),this.arguments.get(1),154,this.text);
	}
	@Override
	public Touch getTouch(int x, int y) {
		if(x > this.arguments.get(0) - w/2 && x < this.arguments.get(0) + w/2 &&
		   y > this.arguments.get(1) - h/2 && y < this.arguments.get(1) + h/2)
			return new Touch(this);
		return null;
	}
	@Override
	public Point getPointTouch(int index) {
		return null;
	}

}
