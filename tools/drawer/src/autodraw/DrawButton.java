package autodraw;

import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics2D;

public class DrawButton {
	private final int position[];
	private final Element.ElementType type;
	public static final int toolButtonWidth = 30;
	public static final int toolButtonHeight = 30;
	private static final int padding = 3;
	
	public DrawButton(int x, int y, Element.ElementType type) {
		this.position = new int[2];
		this.position[0] = x;
		this.position[1] = y;
		this.type = type;
	}
	public int getX() {return this.position[0];}
	public int getY() {return this.position[1];}
	public Element.ElementType getType() {return this.type;}
	public boolean clicked(int x, int y) {
		return x > getX() && x < getX()+toolButtonWidth && y > getY() && getY() < y+toolButtonHeight;
	}
	public void draw(Graphics2D g2d) {
		if(this.type == Element.ElementType.LINE) {
			drawLineButton(g2d,getX(),getY());
		} else if(this.type == Element.ElementType.RECT) {
			drawRectButton(g2d,getX(),getY());
		} else if(this.type == Element.ElementType.POLYGON) {
			drawPolygonButton(g2d,getX(),getY());
		} else if(this.type == Element.ElementType.OVAL) {
			drawOvalButton(g2d,getX(),getY());
		}
	}

	public void drawLineButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.drawRect(x, y, toolButtonWidth, toolButtonHeight);
		g2d.drawLine(x+padding, y+padding, toolButtonWidth-padding*2, toolButtonHeight-padding*2);
	}

	public void drawRectButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.drawRect(x, y, toolButtonWidth, toolButtonHeight);
		g2d.drawRect(x+padding, y+padding, toolButtonWidth-padding*2, toolButtonHeight-padding*2);
	}

	public void drawPolygonButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,12));
		g2d.drawRect(x, y, toolButtonWidth, toolButtonHeight);
		g2d.drawString("Poly", x+padding, y+toolButtonHeight*2/3);
	}

	public void drawOvalButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.drawRect(x, y, toolButtonWidth, toolButtonHeight);
		g2d.drawOval(x+padding, y+padding, toolButtonWidth-padding*2, toolButtonHeight-padding*2);
	}
	
	public void drawButton(Graphics2D g2d, String name, int x, int y) {
		if(name == "line") {
			drawLineButton(g2d,x,y);
		} else if(name == "rect") {
			drawRectButton(g2d,x,y);
		} else if(name == "polygon") {
			drawPolygonButton(g2d,x,y);
		} else if(name == "oval") {
			drawOvalButton(g2d,x,y);
		}
	}
}
