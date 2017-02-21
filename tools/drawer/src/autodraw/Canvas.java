package autodraw;

import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.Point;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.KeyAdapter;
import java.awt.event.KeyEvent;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;
import java.awt.event.MouseMotionListener;
import java.util.ArrayList;

import javax.swing.JPanel;

public class Canvas extends JPanel implements ActionListener,MouseListener,MouseMotionListener {

	private static final long serialVersionUID = 3196872295964223375L;
	
	private final int toolbarHeight = 30;
	private final int toolButtonWidth = 30;
	private final int padding = 3;
	private final int buttonCount = 4;
	
	private int buttonPositions[][] = {
			{0,0},{toolButtonWidth,0},{toolButtonWidth*2,0},{toolButtonWidth*3,0}	
	};
	private String buttonNames[] = {
			"line","rect","polygon","oval"
	};
	private Element.ElementType tools[] = {
			Element.ElementType.LINE, Element.ElementType.RECT,
			Element.ElementType.POLYGON, Element.ElementType.OVAL
	};
	
	private Element.ElementType tool = Element.ElementType.LINE;
	private ArrayList<Point> pointBuffer;
	private int mousex,mousey;
	
	private ArrayList<Element> elementList;
	
	public Canvas() {
		repaint();
		this.pointBuffer = new ArrayList<Point>();
		this.elementList = new ArrayList<Element>();
		addKeyListener(new KeyAdapter() {
			public void keyPressed(KeyEvent e)
			{
				int key = e.getKeyCode();
				switch(key)
				{
				case KeyEvent.VK_W:
					System.out.println("W");
					break;
				case KeyEvent.VK_S:
					System.out.println("S");
					break;
				}
				repaint();
			}
		});
		addMouseListener(this);
		addMouseMotionListener(this);
	}

	public void actionPerformed(ActionEvent ae) {
		String cmd = ae.getActionCommand();
		if(cmd.equals("")) {
			
		}
		System.out.println("action");
		repaint();
	}
	
	public void paint(Graphics g) {
		Graphics2D g2d = (Graphics2D) g;
		
		g2d.setColor(Color.white);
		g2d.fillRect(0, 0, this.getWidth(), this.getHeight());
		
		if(this.tool == Element.ElementType.LINE) {
			if(this.pointBuffer.size() >= 1) {
				Point startPoint = pointBuffer.get(0);
				g2d.setColor(Color.black);
				g2d.drawLine(startPoint.x,startPoint.y,mousex,mousey);
			}
		} else if (this.tool == Element.ElementType.RECT) {
			if(this.pointBuffer.size() >= 1) {
				Point startPoint = pointBuffer.get(0);
				g2d.setColor(Color.black);
				g2d.drawRect(startPoint.x,startPoint.y,mousex-startPoint.x,mousey-startPoint.y);
			}
		} else if (this.tool == Element.ElementType.OVAL) {
			if(this.pointBuffer.size() >= 1) {
				Point startPoint = pointBuffer.get(0);
				g2d.setColor(Color.black);
				g2d.drawOval(startPoint.x*2-mousex,startPoint.y*2-mousey,
						mousex*2-startPoint.x*2,mousey*2-startPoint.y*2);
			}
		}
		
		for(Element e: this.elementList) {
			e.draw(g2d);
		}
		
		for(int i = 0; i < buttonCount; i++) {
			drawButton(g2d,buttonNames[i],buttonPositions[i][0],buttonPositions[i][1]);
		}
	}

	public void drawLineButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.drawRect(x, y, toolButtonWidth, toolbarHeight);
		g2d.drawLine(x+padding, y+padding, toolButtonWidth-padding*2, toolbarHeight-padding*2);
	}

	public void drawRectButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.drawRect(x, y, toolButtonWidth, toolbarHeight);
		g2d.drawRect(x+padding, y+padding, toolButtonWidth-padding*2, toolbarHeight-padding*2);
	}

	public void drawPolygonButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,12));
		g2d.drawRect(x, y, toolButtonWidth, toolbarHeight);
		g2d.drawString("Poly", x+padding, y+toolbarHeight*2/3);
	}

	public void drawOvalButton(Graphics2D g2d, int x, int y) {
		g2d.setColor(Color.black);
		g2d.drawRect(x, y, toolButtonWidth, toolbarHeight);
		g2d.drawOval(x+padding, y+padding, toolButtonWidth-padding*2, toolbarHeight-padding*2);
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
	
	public boolean buttonClicked(int x, int y, int bx, int by) {
		return x > bx && x < bx+toolButtonWidth && y > by && y < by+toolbarHeight;
	}

	public void mouseDragged(MouseEvent e) {
		// TODO Auto-generated method stub
		
	}

	public void mouseMoved(MouseEvent e) {
		this.mousex = e.getX();
		this.mousey = e.getY();
		repaint();
	}

	public void mouseClicked(MouseEvent e) {
		int x = e.getX(), y = e.getY();
		boolean clickedButton = false;
		this.mousex = x;
		this.mousey = y;
		
		for(int i = 0; i < buttonCount; i++) {
			if(buttonClicked(x,y,buttonPositions[i][0],buttonPositions[i][1])) {
				this.tool = this.tools[i];
				this.pointBuffer = new ArrayList<Point>();
				clickedButton = true;
			}
		}
		if(!clickedButton) {
			this.pointBuffer.add(new Point(x,y));
		}
		if(this.tool == Element.ElementType.LINE) {
			if(this.pointBuffer.size() >= 2) {
				Point p1 = this.pointBuffer.get(0);
				Point p2 = this.pointBuffer.get(1);
				this.elementList.add(new LineElement(p1.x,p1.y,p2.x,p2.y));
				this.pointBuffer = new ArrayList<Point>();
			}
		} else if(this.tool == Element.ElementType.RECT) {
			if(this.pointBuffer.size() >= 2) {
				Point p1 = this.pointBuffer.get(0);
				Point p2 = this.pointBuffer.get(1);
				this.elementList.add(new RectElement(p1.x,p1.y,p2.x,p2.y));
				this.pointBuffer = new ArrayList<Point>();
			}
		} else if(this.tool == Element.ElementType.POLYGON) {
			
		} else if(this.tool == Element.ElementType.OVAL) {
			if(this.pointBuffer.size() >= 2) {
				Point p1 = this.pointBuffer.get(0);
				Point p2 = this.pointBuffer.get(1);
				this.elementList.add(new OvalElement(p1.x,p1.y,p2.x-p1.x,p2.y-p1.y));
				this.pointBuffer = new ArrayList<Point>();
			}
		}
		repaint();
	}

	public void mouseEntered(MouseEvent e) {
		// TODO Auto-generated method stub
		
	}

	public void mouseExited(MouseEvent e) {
		// TODO Auto-generated method stub
		
	}

	public void mousePressed(MouseEvent e) {
		// TODO Auto-generated method stub
		
	}

	public void mouseReleased(MouseEvent e) {
		// TODO Auto-generated method stub
		
	}
}
