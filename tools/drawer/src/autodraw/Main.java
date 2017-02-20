package autodraw;

import java.awt.BorderLayout;
import java.awt.Dimension;
import javax.swing.JFrame;

import autodraw.Canvas;

public class Main extends JFrame {

	private static final long serialVersionUID = -2449636255315974141L;
	private Canvas canvas;

	public Main(String s) {
		super(s);
		this.canvas = new Canvas();
		this.canvas.setFocusable(true);
		this.canvas.setSize(new Dimension(800,600));
		
		add(this.canvas);
		setLayout(new BorderLayout());
		setSize(800,600);
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setResizable(false);
		setVisible(true);
	}

	/**
	 * @param args: command line arguments to this program
	 */
	public static void main(String[] args) {
		new Main("Autodraw");
	}

}
