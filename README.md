# Alkitab API

## Multi-language Bible API

Alkitab API is a RESTful API developed using Golang and MongoDB providing Bible verses in three languages:
- 🇮🇩 Indonesian
- 🇻🇦 Latin
- 🇹🇴 Toraja

This API is designed for Bible study applications, cross-language learning, and digital projects requiring structured access to scripture data.

## Technology Stack

- [Golang](https://golang.org/) - Backend programming language
- [MongoDB](https://www.mongodb.com/) - Database system
- RESTful architecture
- JSON response format

## Data Structure

### Document Format

```json
{
  "title": "Matthew",
  "book_number": "1",
  "chapter": "1",
  "verses": [
    {
      "verse_number": 1,
      "content": {
        "Latin": "Liber generationis Iesu Christi filii David filii Abraham.",
        "Indonesia": "Inilah silsilah Yesus Kristus, anak Daud, anak Abraham.",
        "Toraja": "Inde sia tu ossoran nene'Na Yesu Kristus, anakna Daud, anakna Abraham."
      }
    },
    {
      "verse_number": 2,
      "content": {
        "Latin": "Abraham genuit Isaac...",
        "Indonesia": "Abraham memperanakkan Ishak...",
        "Toraja": "Abraham menting anakna Isaac..."
      }
    }
  ]
}
```

## MongoDB Structure

- Main collection: `new-testament`
- Each document represents one Bible chapter

## Scripture Source References

| Language   | Source                                                       |
|------------|--------------------------------------------------------------|
| Latin      | [biblestudytools.com](https://www.biblestudytools.com/)      |
| Indonesian | [imankatolik.or.id](https://www.imankatolik.or.id/alkitab.php) |
| Toraja     | [alkitab.mobi](https://alkitab.mobi/)                        |

## Implementation Notes

- This API was developed as a study and ministry project
- Data is managed, verified, and expanded manually by contributors
- All content undergoes validation to ensure accuracy and precision

## Contributing

We welcome contributions from developers! If you would like to add translations or improve existing data, please submit a pull request or open an issue in this repository.

## License

This project is open-source and freely available for educational purposes, ministry work, and personal development.