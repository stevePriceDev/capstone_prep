const mongoose = require('mongoose');

if (process.argv.length < 3) {
  console.log('give password as argument');
  process.exit(1);
}

const password = process.argv[2];

const url = 
  `mongodb+srv://sgp1986:${password}@cluster0.zgyneh9.mongodb.net/testNoteApp?retryWrites=true&w=majority`;

mongoose.set('strictQuery', false);
mongoose.connect(url);

const noteSchema = new mongoose.Schema({
  content: String,
  important: Boolean,
});

const Note = mongoose.model('Note', noteSchema);

// const note = new Note({
//   content: 'HTML is Easy',
//   important: true,
// });

const note = new Note({
  "id": 2,
  "content": "Browser can execute only JavaScript",
  "important": true
});

// const note = new Note({
//   "id": 3,
//   "content": "GET and POST are the most important methods of HTTP protocol",
//   "important": false
// });

note.save().then(result => {
  console.log('note saved!');
  mongoose.connection.close();
});

// Note.find({}).then(result => {
//   result.forEach(note => {
//     console.log(note);
//   });
//   mongoose.connection.close();
// });
