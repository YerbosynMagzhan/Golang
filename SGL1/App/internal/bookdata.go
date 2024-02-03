package internal

type Book struct {
	Title       string
	Genre       string
	Author      string
	Description string
	Characters  []Character
}

type Character struct {
	Name  string
	Role  string
	Image string
}

func GetBookData() Book {
	return Book{
		Title:       "Wuthering Heights",
		Genre:       "novel",
		Author:      "Emily Bronte",
		Description: "Wuthering Heights is the first and only novel by the English author Emily Brontë, initially published in 1847 under her pen name 'Ellis Bell'. It concerns two families of the landed gentry living on the West Yorkshire moors, the Earnshaws and the Lintons, and their turbulent relationships with the Earnshaws' foster son, Heathcliff. The novel was influenced by Romanticism and Gothic fiction.",
		Characters: []Character{
			{"Heathcliff", "Protagonist", "ui/static/i.webp"},
			{"Catherine Earnshaw", "Heroine", "ui/static/i.webp"},
		},
	}
}

func GetCharacterByName(name string) (Character, bool) {
	for _, char := range GetBookData().Characters {
		if char.Name == name {
			return char, true
		}
	}
	return Character{}, false
}
