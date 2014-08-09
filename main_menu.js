var MainMenu = {
	Paint: function(offsetX, offsetY, left, top, width, height) {
		this.lastWidth = width;
		var optionTop = this.optionTop + offsetY;
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

	HandleKey: function(key) {
		switch (key) {
		case 'MenuPrev':
			this.selection += this.options.length - 1;
			this.selection %= this.options.length;
			Repaint();
			return true;

		case 'MenuNext':
			this.selection += 1;
			this.selection %= this.options.length;
			Repaint();
			return true;

		case 'MenuCancel':
			Exit();
			return true;

		case 'MenuSelect':
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

	HandleRune: function(ch) {
		return false;
	},

	HandleMouse: function(x, y) {
		var option = y - this.optionTop;
		if (0 <= option && option < this.options.length) {
			var optionLeft = Math.floor((this.lastWidth - this.options[option].length) / 2);
			if (optionLeft <= x && x < optionLeft + this.options[option].length) {
				if (option == this.selection) {
					return this.HandleKey(termbox.KeyEnter, 0);
				} else {
					this.selection = option;
					Repaint();
					return true;
				}
			}
		}
		return false;
	},

	lastWidth: 1,
	selection: 0,
	options: ['New world'.split(''), 'Exit'.split('')],
	optionTop: 2
};

var border = new BorderPanel();
border.SetFg(termbox.ColorBlack);
border.SetBg(termbox.ColorWhite);
border.SetTitle("Suudsu");

border.SetChild(MainMenu);
UI().SetChild(border);

Repaint();
