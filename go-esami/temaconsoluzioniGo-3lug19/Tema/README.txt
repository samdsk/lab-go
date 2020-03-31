====== Istruzioni per l'esame ======

===== Scompattare i file =====

Con 'tar xvf Tema.tar' viene espanso l'archivio a partire dalla directory in cui ci si trova.

In questa si troverà quindi una serie di sottodirectory ognuna delle quali conterrà un file nella forma '<nome esercizio>_test.go'.

In testa a questi file si trova, sotto forma di commento, il testo dell'esercizio. Di seguito si trovano i test da eseguire per fare una prima (auto)valutazione del proprio elaborato.



===== Uso dei test =====

Per svolgere il compito bisogna entrare in ogni sottodirectory e creare il relativo file 'go' (se ad esempio il test si chiama 'filtro_test.go' va creato il file 'filtro.go', attenzione al case, in generale sono caratteri tutti minuscoli) secondo le specifiche.

Per "testare" il codice basta compilare innanzitutto il proprio elaborato con 'go build', poi si può lanciare il testing con 'go test -v'. Ricordarsi di ricompilare dopo ogni modifica prima di lanciare di nuovo i test.

PASS indica che il programma x.go ha superato i test contenuti in x_test.go (ma tra questi potrebbero esserci test che fanno vedere output ottenuto e/o output atteso, senza fare nessun confronto né controllo, per i quali quindi PASS indica solo un superamento formale e non sui risultati)
FAIL indica che il programma x.go NON ha superato almeno uno dei test contenuti in x_test.go

Si consiglia VIVAMENTE di esaminare anche il codice dei test oltre al testo dell'esercizio al fine di comprendere meglio ciò che viene chiesto per lo svolgimento.

Nota Bene: i test effettivi eseguiti dai docenti in fase di correzione potrebbero essere in numero e tipo diversi da quelli presentati nel tema d'esame. I.e., un PASS non significa automaticamente che l'esercizio si privo di errori. Inoltre a determinare il voto concorreranno anche altri aspetti, quali la semplicità del codice, la leggibilità, l'uso della memoria, ecc.



===== Consegna =====

Va effettuata caricando i SINGOLI file '<nome esercizio>.go' (quindi NON vanno consegnati i test) su https://upload.di.unimi.it utilizzando la propria login e scegliendo la sessione corretta (nel dubbio chiedere ai docenti).

Si può caricare più volte lo stesso file, verrà valutata soltanto l'ultima versione depositata. Si consiglia di caricare il file anche 'in itinere' in modo da avere un backup in caso di problemi alla postazione.

Nota bene: la sessione (login) di upload scade dopo circa 15 minuti di inattività, quindi nel caso occorre semplicemente fornire di nuovo le proprie credenziali (evitare trovarsi in questa situazione a un minuto dall'ora di consegna).

***ATTENZIONE***: la sessione d'esame si CHIUDE AUTOMATICAMENTE all'orario "di consegna" che viene comunicato in aula dai docenti. Fino a quell'ora si possono consegnare i propri elaborati, oltre no e NON SI FARANNO ECCEZIONI. Al termine della consegna il sistema di upload NON ACCETTA più il caricamento dei file. Si consiglia di non ridursi all'ultimo momento.



===== Per ritirarsi =====

Occorre caricare un file vuoto dal nome 'ritirato.txt' (entro l'orario di consegna).



===== Valutazione =====

***ATTENZIONE***: l'esercizio 'filtro' deve compilare e girare correttamente (rispetto alle specifiche), ciò è condizione necessaria affinché tutto l'elaborato venga valutato.

Verranno valutati (cioè ne verrà esaminato e giudicato il sorgente) solo gli esercizi (compreso il filtro) che compilano senza errori, gli altri non verranno presi in considerazione ai fini del voto.

Gli esercizi hanno pesi diversi per cui non c'è un "numero minimo di esercizi giusti per passare l'esame". I pesi vengono assegnati a valle di una valutazione sul campo della difficoltà di svolgimento, quindi chiedere "quali sono i pesi?" durante l'esame non può ricevere risposta.



===== Materiale utilizzabile =====

* libro di testo
* documentazione online di Go
* penna
* carta fornita dai docenti



===== Materiale NON utilizzabile =====

Ogni altro oggetto non elencato nella sezione precedente. Quindi, ad esempio, NON si possono tenere in vista telefoni, appunti, altri libri, quaderni, etc.

Va riposto tutto nel proprio zaino/borsa che poi andrà deposto a terra su un lato dell'aula (seguire indicazioni del docente).



===== Comportamenti sanzionabili =====

Parlare con altri studenti a esame iniziato comporta espulsione e annullamento dell'esame, per TUTTI i soggetti coinvolti nella comunicazione. Se si ha bisogno di delucidazioni sul tema d'esame o per altre domande rivolgersi (per alzata di mano o avvicinandosi alla cattedra) ai docenti.

Idem in caso di utilizzo di materiale non consentito.