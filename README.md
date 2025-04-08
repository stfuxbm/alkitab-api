# Alkitab API

A RESTful API providing Bible chapters in three languages:

- 🇮🇩 Indonesian  
- 🇻🇦 Latin  
- 🇹🇴 Toraja  

Built with Golang and MongoDB, this API is ideal for Bible study, language comparison, and digital scripture applications.

---

## Tech Stack

- **Golang** — for backend API  
- **MongoDB** — NoSQL database  
- **JSON** — response format  

---

## Data Structure

Each document in the `new-testament` collection represents one chapter of the Bible:

```json
{
  "title": "Matthew",
  "book_number": "1",
  "chapter": "1",
  "verses": [
    {
      "verse_number": 1,
      "content": {
        "Latin": "Liber generationis Iesu Christi...",
        "Indonesia": "Inilah silsilah Yesus Kristus...",
        "Toraja": "Inde sia tu ossoran nene'Na Yesu Kristus..."
      }
    }
  ],
  "metadata": {
    "latin": {
      "language": "Latin",
      "version": "Biblia Sacra Vulgata",
      "publisher": "vulgate.org"
    },
    "indonesia": {
      "language": "Indonesia",
      "version": "Iman Katolik",
      "publisher": "imankatolik.or.id"
    },
    "toraja": {
      "language": "Toraja",
      "version": "SDA Toraja",
      "publisher": "alkitab.me"
    }
  },
  "createdat": "timestamp",
  "updatedat": "timestamp"
}
```

---

## Source References

| Language   | Source                       |
|------------|------------------------------|
| Indonesian | [imankatolik.or.id](https://www.imankatolik.or.id/alkitab.php) |
| Latin      | [vulgate.org](https://vulgate.org/)                            |
| Toraja     | [alkitab.me](https://alkitab.me/sda-toraja)                   |

---

## Notes

- Each `title` + `chapter` combination is unique.
- Metadata is validated per language entry.
- Data is manually curated and reviewed.

---

## Contributing

Contributions are welcome!  
Open an issue or submit a pull request to improve translations, features, or documentation.

---

## License

This project is licensed under the [MIT License](LICENSE).  
You are free to use, modify, and distribute this project for educational, ministry, or personal development purposes.

---