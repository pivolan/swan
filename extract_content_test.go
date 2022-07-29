package swan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromHTML1(t *testing.T) {
	a, e := FromHTML("", []byte(`<p>von Frank Müller</p>
<p>Dieser Artikel ist ein Test für das Ticket CUE-1257.</p>
<p>Dies ist eine Änderung zu
    Testzwecken von Michael Karl dem großen Helden. Einer wie keiner, er nimmt es den Reichen und schenkt es den Armen.
    Im Folgenden Text werden einmal sämtliche Elemente getestet, die für Text-to-Speech relevant sind. Es folgt die
    Einbettung von Artikelempfehlungen bzw. Related Content: Es folgen die Überschriften. Hier ist ein bisschen Text zum
    Artikel. Hier ist noch mehr Text zum Artikel. Es folgt eine Tabelle: An dieser Stelle finden Sie im Web oder in der
    NewsApp eine Tabelle zum Thema. Es folgt ein Video: An dieser Stelle können Sie im Web oder der NewsApp ein Video
    ansehen. Es folgt eine eingebettete Bildergalerie: Es folgt ein Livestream: Es folgt eine Google-Karte: Es folgt ein
    PDF: Es folgt ein Zeitstrahl: Es folgt ein Cockpit-Liveticker per oEmbed: Es folgt ein Inline-Bild: Es folgt eine
    nummerierte Liste:
<ol>
    <li>Dies ist der erste Punkt</li>
    <li>Dies ist der zweite Punkt</li>
    <li>Dies ist der dritte Punkt</li>
    <li>Dies ist der vierte Punkt</li>
</ol>

Es folgt eine unnummerierte Liste:
<ul>
    <li>Erster Punkt</li>
    <li>Zweiter Punkt</li>
    <li>Dritter Punkt</li>
    <li>Vierter Punkt</li>
</ul>
Es folgt ein Infokasten: Weitere Informationen zum Thema Infobox Text macht keinen Sinn, ist aber spannend zu hören.
Es folgt ein Zitatkasten: Ringdingdingding Ringdingdingding so Fuchs (Tier des Waldes) Es folgt ein redaktioneller Kommentar: Es folgt Facebook-Post per oEmbed: Es folgt ein Twitter-Post per oEmbed: Es folgt ein Instagram-Post per oEmbed: Es folgt ein youtube-Video per oEmbed: Es folgt ein IFrame: An dieser Stelle ist ein externer Inhalt eingebunden. Es folgt eine HTML-Box: An dieser Stelle ist ein externer Inhalt eingebunden. Dies ist das Ende des Textes
</p>`))
	fmt.Println("CLEANED TEXT: ", a.CleanedText, e)
}
func TestPyContentExtractors(t *testing.T) {
	t.Parallel()

	runPyTests(t,
		"test_data/python-goose/content/",
		func(t *testing.T, name string, a *Article, r *Result) {
			e := r.Expected

			if e.MetaDescription != "" && e.MetaDescription != a.Meta.Description {
				t.Fatalf(
					"%s: MetaDescription does not match:\n"+
						"	Got: %s\n"+
						"	Expected: %s",
					name, a.Meta.Description, e.MetaDescription)
			}

			if e.MetaKeywords != "" && e.MetaKeywords != a.Meta.Keywords {
				t.Fatalf(
					"%s: MetaKeywords does not match:\n"+
						"	Got: %s\n"+
						"	Expected: %s",
					name, a.Meta.Keywords, e.MetaKeywords)
			}

			if e.Title != "" && e.Title != a.Meta.Title {
				t.Fatalf(
					"%s: Title does not match:\n"+
						"	Got: %s\n"+
						"	Expected: %s",
					name, a.Meta.Title, e.Title)
			}

			if e.MetaLang != "" && e.MetaLang != a.Meta.Lang {
				t.Fatalf(
					"%s: Lang does not match:\n"+
						"	Got: %s\n"+
						"	Expected: %s",
					name, a.Meta.Lang, e.MetaLang)
			}

			cleaned := a.CleanedText
			if len(r.Expected.CleanedText) < len(cleaned) {
				cleaned = cleaned[:len(r.Expected.CleanedText)]
			}

			assert.Equal(t, r.Expected.CleanedText, cleaned)
			if cleaned != r.Expected.CleanedText {
				t.Fatalf(
					"%s: CleanedText does not match:\n"+
						"	Got:      %s\n"+
						"	Expected: %s",
					name, cleaned, r.Expected.CleanedText)
			}
		})
}
