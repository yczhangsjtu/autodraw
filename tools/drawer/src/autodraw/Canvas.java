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
	private final int buttonCount = 5;
	
	private DrawButton buttons[];
	
	private Element.ElementType tool = Element.ElementType.LINE;
	private ArrayList<Point> pointBuffer;
	private String textBuffer;
	private int mousex,mousey;
	
	private ArrayList<Element> elementList;
	private int originx,originy;
	private JTextArea output;

	private int grid = 1;

	public Canvas() {
		repaint();
		this.pointBuffer = new ArrayList<Point>();
		this.elementList = new ArrayList<Element>();
		this.textBuffer = "";
		this.buttons = new DrawButton[buttonCount];
		this.buttons[0] = new DrawButton(0,0,Element.ElementType.LINE);
		this.buttons[1] = new DrawButton(DrawButton.toolButtonWidth,0,Element.ElementType.RECT);
		this.buttons[2] = new DrawButton(DrawButton.toolButtonWidth*2,0,Element.ElementType.POLYGON);
		this.buttons[3] = new DrawButton(DrawButton.toolButtonWidth*3,0,Element.ElementType.OVAL);
		this.buttons[4] = new DrawButton(DrawButton.toolButtonWidth*4,0,Element.ElementType.TEXT);
		addKeyListener(this);
		addMouseListener(this);
		addMouseMotionListener(this);
	}
	
	public void paint(Graphics g) {
		Graphics2D g2d = (Graphics2D) g;
		g2d.setColor(Color.white);
		g2d.fillRect(0, 0, this.getWidth(), this.getHeight());
		drawAxis(g2d);
		drawTmp(g2d);
		drawElements(g2d);
		drawButtons(g2d);
		drawCoordinates(g2d);
		drawMouse(g2d);
	}
	

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
	
	private void drawElements(Graphics2D g2d) {
		for(Element e: this.elementList)
			e.draw(g2d);
	}
	
	private void drawButtons(Graphics2D g2d) {
		for(int i = 0; i < buttonCount; i++) 
			buttons[i].draw(g2d);
	}
	
	private void drawCoordinates(Graphics2D g2d) {
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,12));
		g2d.drawString(String.format("%d,%d", this.mousex-this.originx, -(this.mousey-this.originy)),
				getWidth()-100, toolbarHeight*2/3);
	}
	
	private void drawAxis(Graphics2D g2d) {
		g2d.setColor(Color.green);
		g2d.drawLine(0,this.originy,getWidth(),this.originy);
		g2d.drawLine(this.originx,0,this.originx,getHeight());
	}
	
	private void drawTmp(Graphics2D g2d) {
		if(this.tool == Element.ElementType.LINE) {
			drawTmpLine(g2d);
		} else if (this.tool == Element.ElementType.RECT) {
			drawTmpRect(g2d);
		} else if (this.tool == Element.ElementType.OVAL) {
			drawTmpOval(g2d);
		} else if (this.tool == Element.ElementType.POLYGON) {
			drawTmpPolygon(g2d);
		} else if (this.tool == Element.ElementType.TEXT) {
			drawTmpText(g2d);
		}
	}
	
	private void drawTmpLine(Graphics2D g2d) {
		if(this.pointBuffer.size() >= 1) {
			Point startPoint = pointBuffer.get(0);
			g2d.setColor(Color.black);
			g2d.drawLine(startPoint.x,startPoint.y,this.mousex,this.mousey);
		}
	}
	
	private void drawTmpRect(Graphics2D g2d) {
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
	}

	private void drawTmpPolygon(Graphics2D g2d) {
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

	private void drawTmpOval(Graphics2D g2d) {
		if(this.pointBuffer.size() >= 1) {
			Point startPoint = pointBuffer.get(0);
			g2d.setColor(Color.black);
			int a = Math.abs(this.mousex-startPoint.x), b = Math.abs(this.mousey-startPoint.y);
			g2d.drawOval(startPoint.x-a,startPoint.y-b,a*2,b*2);
		}
	}
	
	private void drawTmpText(Graphics2D g2d) {
		g2d.setFont(new Font("TimesRoman",Font.PLAIN,24));
		if(this.pointBuffer.size() >= 1) {
			Point point = pointBuffer.get(0);
			g2d.setColor(Color.black);
			int w = g2d.getFontMetrics().stringWidth(this.textBuffer);
			int h = g2d.getFontMetrics().getHeight();
			g2d.drawString(textBuffer, point.x-w/2, point.y+h/2);
			g2d.drawLine(point.x+w/2, point.y-h/2, point.x+w/2, point.y+h/2);
		}
	}

	private void drawMouse(Graphics2D g2d) {
		g2d.drawOval(this.mousex-5, this.mousey-5, 10, 10);
		g2d.drawOval(this.mousex-1, this.mousey-1, 1, 1);
	}

	public void mouseDragged(MouseEvent e) {

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
		this.mousex = e.getX();
		this.mousey = e.getY();
		rectifyMouse();
		if(!this.isFocusOwner()) this.grabFocus();
		
		if(e.getButton() == MouseEvent.BUTTON1) {
			if(!clickButton(this.mousex,this.mousey)) {
				this.pointBuffer.add(new Point(this.mousex,this.mousey));
				completeFixed();
			}
		} else if(e.getButton() == MouseEvent.BUTTON2) {
			clear();
		} else if(e.getButton() == MouseEvent.BUTTON3) {
			if(this.tool == Element.ElementType.POLYGON) {
				completePolygon();
			} else {
				clear();
			}
		}
		repaint();
	}
	
	public boolean clickButton(int x, int y) {
		for(int i = 0; i < buttonCount; i++) {
			if(buttons[i].clicked(x, y)) {
				this.tool = buttons[i].getType();
				this.pointBuffer = new ArrayList<Point>();
				return true;
			}
		}
		return false;
	}
	
	public void clear() {
		this.pointBuffer = new ArrayList<Point>();
		this.textBuffer = "";
	}
	
	public void completeFixed() {
		if(this.tool == Element.ElementType.LINE) {
			completeLine();
		} else if(this.tool == Element.ElementType.RECT) {
			completeRect();
		} else if(this.tool == Element.ElementType.OVAL) {
			completeOval();
		}
	}
	
	public void completeLine() {
		if(this.pointBuffer.size() >= 2) {
			Point p1 = this.pointBuffer.get(0);
			Point p2 = this.pointBuffer.get(1);
			this.elementList.add(new LineElement(p1.x,p1.y,p2.x,p2.y));
			clear();
		}
	}
	
	public void completeRect() {
		if(this.pointBuffer.size() >= 2) {
			Point p1 = this.pointBuffer.get(0);
			Point p2 = this.pointBuffer.get(1);
			this.elementList.add(new RectElement(p1.x,p1.y,p2.x,p2.y));
			clear();
		}
	}
	
	public void completeOval() {
		if(this.pointBuffer.size() >= 2) {
			Point p1 = this.pointBuffer.get(0);
			Point p2 = this.pointBuffer.get(1);
			this.elementList.add(new OvalElement(p1.x,p1.y,p2.x-p1.x,p2.y-p1.y));
			clear();
		}
	}
	
	public void completePolygon() {
		if(this.pointBuffer.size() >= 2) {
			this.elementList.add(new PolygonElement(this.pointBuffer));
			clear();
		}
	}
	
	public void completeText() {
		if(this.pointBuffer.size() >= 1 && !textBuffer.equals("")) {
			Point p = this.pointBuffer.get(0);
			this.elementList.add(new TextElement(textBuffer,p.x,p.y));
			clear();
		}
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
		if(this.tool == Element.ElementType.TEXT && this.pointBuffer.size() >= 1) {
			char c = e.getKeyChar();
			if(key == KeyEvent.VK_ENTER) {
				completeText();
			} else if(key == KeyEvent.VK_BACK_SPACE) {
				if(this.textBuffer.length() > 0)
					this.textBuffer = this.textBuffer.substring(0, this.textBuffer.length()-1);
			} else if(c >= ' ' && c <= '~')
				this.textBuffer += c;
		} else {
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
				break;
			}
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
