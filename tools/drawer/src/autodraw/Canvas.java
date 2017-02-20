package autodraw;

import java.awt.Color;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JPanel;

public class Canvas extends JPanel implements ActionListener {

	private static final long serialVersionUID = 3196872295964223375L;
	
	public Canvas() {
		repaint();
	}

	public void actionPerformed(ActionEvent ae) {
		String cmd = ae.getActionCommand();
		if(cmd.equals("")) {
			
		}
		repaint();
	}
	
	public void paint(Graphics g) {
		Graphics2D g2d = (Graphics2D) g;
		
		g2d.setColor(Color.black);
		g2d.drawRect(0, this.getWidth(), 0, this.getHeight());
		g2d.drawRect(100, this.getWidth(), 100, this.getHeight());
		System.out.println("Rect");
	}

}
