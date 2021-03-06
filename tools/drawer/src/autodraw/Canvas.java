package autodraw;

import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.Point;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;
import java.awt.event.MouseMotionListener;
import java.util.ArrayList;

import javax.swing.JPanel;
import javax.swing.JTextArea;

public class Canvas extends JPanel implements MouseListener,MouseMotionListener,KeyListener {

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
	private int originx,originy;
	private JTextArea output;

	private int grid = 1;
	
	public JTextArea getOutput() {
		return output;
	}

	public void setOutput(JTextArea output) {
		this.output = output;
	}

	public int getOriginy() {
		return originy;
	}

	public void setOriginy(int originy) {
		this.originy = originy;
	}

	public int getOriginx() {
		return originx;
	}

	public void setOriginx(int originx) {
		this.originx = originx;
	}

	public Canvas() {
		repaint();
		this.pointBuffer = new ArrayList<Point>();
		this.elementList = new ArrayList<Element>();
		addKeyListener(this);
		addMouseListener(this);
		addMouseMotionListener(this);
	}
	
	public void paint(Graphics g) {
		Graphics2D g2d = (Graphics2D) g;
		
		g2d.setColor(Color.white);
		g2d.fillRect(0, 0, this.getWidth(), this.getHeight());
		
		g2d.setColor(Color.green);
		g2d.drawLine(0,this.originy,getWidth(),this.originy);
		g2d.drawLine(this.originx,0,this.originx,getHeight());
		
		if(this.tool == Element.ElementType.LINE) {
			if(this.pointBuffer.size() >= 1) {
				Point startPoint = pointBuffer.get(0);
				g2d.setColor(Color.black);
				g2d.drawLine(startPoint.x,startPoint.y,this.mousex,this.mousey);
			}
		} else if (this.tool == Element.ElementType.RECT) {
			if(this.pointBuffer.size() >= 1) {
				Point startPoint = pointBuffer.get(0);
				g2d.setColor(Color.black);
				int x0 = startPoint.x, y0 = startPoint.y, x1 = this.mousex, y1 = this.mousey;
				if(x0 > x1) {
					int tmp = x0; x0 = x1; x1 = tmp;
				}
				if(y0 > y1) {
					int tmp = y0; y0 = y1; y1 = tmp;
				}
				g2d.drawRect(x0,y0,x1-x0,y1-y0);
			}
		} else if (this.tool == Element.ElementType.OVAL) {
			if(this.pointBuffer.size() >= 1) {
				Point startPoint = pointBuffer.get(0);
				g2d.setColor(Color.black);
				int a = Math.abs(this.mousex-startPoint.x), b = Math.abs(this.mousey-startPoint.y);
				g2d.drawOval(startPoint.x-a,startPoint.y-b,a*2,b*2);
			}
		} else if (this.tool == Element.ElementType.POLYGON) {
			if(this.pointBuffer.size() >= 1) {
				g2d.setColor(Color.black);
				for(int i = 0; i < this.pointBuffer.size()-1; i++) {
					Point p1 = this.pointBuffer.get(i);
					Point p2 = this.pointBuffer.get(i+1);
					g2d.drawLine(p1.x, p1.y, p2.x, p2.y);
				}
				Point p1 = this.pointBuffer.get(this.pointBuffer.size()-1);
				g2d.drawLine(p1.x, p1.y, mousex, mousey);
			}
		}
		
		for(Element e: this.elementList) {
			e.draw(g2d);
		}
		
		for(int i = 0; i < buttonCount; i++) {
			drawButton(g2d,buttonNames[i],buttonPositions[i][0],buttonPositions[i][1]);
		}
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,12));
		g2d.drawString(String.format("%d,%d", this.mousex-this.originx, -(this.mousey-this.originy)),
				getWidth()-100, toolbarHeight*2/3);
		drawMouse(g2d);
	}

	private void drawMouse(Graphics2D g2d) {
		g2d.drawOval(this.mousex-5, this.mousey-5, 10, 10);
		g2d.drawOval(this.mousex-1, this.mousey-1, 1, 1);
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
		rectifyMouse();
		if(!this.isFocusOwner()) this.grabFocus();
		repaint();
	}

	private void rectifyMouse() {
		this.mousex = (int)Math.round((double)(this.mousex-this.originx)/this.grid)*this.grid+this.originx;
		this.mousey = (int)Math.round((double)(this.mousey-this.originy)/this.grid)*this.grid+this.originy;;
	}

	public void mouseClicked(MouseEvent e) {
		int x = e.getX(), y = e.getY();
		boolean clickedButton = false;
		this.mousex = x;
		this.mousey = y;
		rectifyMouse();
		x = this.mousex;
		y = this.mousey;
		if(!this.isFocusOwner()) this.grabFocus();
		
		if(e.getButton() == MouseEvent.BUTTON1) {
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
			} else if(this.tool == Element.ElementType.OVAL) {
				if(this.pointBuffer.size() >= 2) {
					Point p1 = this.pointBuffer.get(0);
					Point p2 = this.pointBuffer.get(1);
					this.elementList.add(new OvalElement(p1.x,p1.y,p2.x-p1.x,p2.y-p1.y));
					this.pointBuffer = new ArrayList<Point>();
				}
			}
		} else if(e.getButton() == MouseEvent.BUTTON2) {
			this.pointBuffer = new ArrayList<Point>();
		} else if(e.getButton() == MouseEvent.BUTTON3) {
			if(this.tool == Element.ElementType.POLYGON) {
				if(this.pointBuffer.size() >= 2) {
					this.elementList.add(new PolygonElement(this.pointBuffer));
					this.pointBuffer = new ArrayList<Point>();
				}
			} else {
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

	public void keyPressed(KeyEvent e) {
		int key = e.getKeyCode();
		switch(key) {
		case KeyEvent.VK_S:
			this.output.setText("");
			for(Element elem: this.elementList) {
				try {
					this.output.setText(this.output.getText()+
							elem.translated(this.originx, this.originy).toString()+"\n");
				} catch (CloneNotSupportedException e1) {
					e1.printStackTrace();
				}
			}
			break;
		case KeyEvent.VK_C:
			this.elementList.clear();
			break;
		case KeyEvent.VK_Z:
			if(this.elementList.size()>0)
				this.elementList.remove(this.elementList.size()-1);
			break;
		case KeyEvent.VK_1:
		case KeyEvent.VK_2:
		case KeyEvent.VK_3:
		case KeyEvent.VK_4:
		case KeyEvent.VK_5:
		case KeyEvent.VK_6:
		case KeyEvent.VK_7:
		case KeyEvent.VK_8:
		case KeyEvent.VK_9:
			this.grid = e.getKeyChar()-'0';
		}
		repaint();
	}

	public void keyReleased(KeyEvent arg0) {
		// TODO Auto-generated method stub
		
	}

	public void keyTyped(KeyEvent arg0) {
		// TODO Auto-generated method stub
		
	}
}
