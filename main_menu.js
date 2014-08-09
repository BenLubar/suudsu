var MainMenu = {
	Paint: function(offsetX, offsetY, left, top, width, height) {
		var optionTop = 2 + offsetY;
		for (var y = 0; y < height; y++) {
			var option = [];
			var optionLeft = 0;
			var color = termbox.ColorWhite;
			if (optionTop <= y && y < optionTop + this.options.length) {
				option = this.options[y - optionTop];
				optionLeft = Math.floor((width - option.length) / 2);
				if (y - optionTop == this.selection) {
					color |= termbox.AttrBold;
				}
			}
			for (var x = 0; x < width; x++) {
				if (optionLeft <= x && x < optionLeft + option.length) {
					termbox.SetCell(left + x, top + y, option[x - optionLeft], color, termbox.ColorBlack);
				}
			}
		}
	},

	HandleKey: function(key, mod) {
		switch (key) {
		case termbox.KeyArrowUp:
			this.selection += this.options.length - 1;
			this.selection %= this.options.length;
			Repaint();
			return true;

		case termbox.KeyArrowDown:
			this.selection += 1;
			this.selection %= this.options.length;
			Repaint();
			return true;

		case termbox.KeyEsc:
			Exit();
			return true;

		case termbox.KeyEnter:
			switch (this.selection) {
			case 0:
				// TODO
				return true;
			case 1:
				Exit();
				return true;
			}
			return false;
		}

		return false;
	},

	HandleRune: function(ch, mod) {
		
	},

	HandleMouse: function(x, y) {
		
	},

	selection: 0,
	options: ['New world'.split(''), 'Exit'.split('')]
};

var border = new BorderPanel();
border.SetFg(termbox.ColorBlack);
border.SetBg(termbox.ColorWhite);
border.SetTitle("Suudsu");

border.SetChild(MainMenu);
UI().SetChild(border);

Repaint();
