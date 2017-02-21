package autodraw;

import java.awt.BorderLayout;
import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.TextArea;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;

import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.JScrollPane;
import javax.swing.JTextArea;

import autodraw.Canvas;

public class Main extends JFrame implements KeyListener {

	private static final long serialVersionUID = -2449636255315974141L;
	private Canvas canvas;
	private JTextArea text;

	public Main(String s) {
		super(s);
		this.canvas = new Canvas();
		this.canvas.setFocusable(true);
		this.canvas.setSize(new Dimension(800,600));
		this.canvas.setLocation(0,0);
		this.canvas.setOriginx(this.canvas.getWidth()/2);
		this.canvas.setOriginy(this.canvas.getHeight()/2);
		this.canvas.addKeyListener(this);
		
		this.text = new JTextArea(10,30);
		this.text.setEditable(false);
		this.canvas.setOutput(this.text);
		JScrollPane scrollPane = new JScrollPane(this.text);
		scrollPane.setSize(new Dimension(200,600));
		JPanel rightPanel = new JPanel();
		rightPanel.setLayout(new BorderLayout());
		rightPanel.add(scrollPane);
		rightPanel.setSize(new Dimension(200,600));

		setLayout(new BorderLayout());
		add(rightPanel,BorderLayout.EAST);
		add(this.canvas,BorderLayout.CENTER);
		setSize(1100,600);
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

    public void paint(Graphics g)
    {
        super.paint(g);
    }

	public void keyPressed(KeyEvent e) {
		switch(e.getKeyCode()) {
		case KeyEvent.VK_ESCAPE:
		case KeyEvent.VK_Q:
			dispose();
		}
	}

	public void keyReleased(KeyEvent arg0) {
		// TODO Auto-generated method stub
		
	}

	public void keyTyped(KeyEvent arg0) {
		// TODO Auto-generated method stub
		
	}
}
