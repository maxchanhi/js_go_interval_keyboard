const keys_obj = {"C Major":["None"],"G Major":["F#"],"F Major":["Bb"],"D Major":["F#","C#"]};

function find_show_key(input_key) {
    return keys_obj[input_key];
}
function showKeySignature() {
    const selectedKey = document.getElementById("keySelect").value;
    const signature = find_show_key(selectedKey);
    document.getElementById("result").textContent = `Key signature for ${selectedKey}: ${signature.join(", ")}`;
}


function list_options() {
    let options = "";
    for (let key in keys_obj) {
      options += `<option value="${key}">${key}</option>`;
    }
    return options;
  }

const notes = ['C', 'C#\nDb', 'D', 'D#\nEb', 'E', 'F', 'F#\nGb', 'G', 
    'G#\nAb', 'A', 'A#\nBb', 'B',"C'", "C#\nDb'", "D'", "D#\nEb'",
    "E'", "F'", "F#\nGb'", "G'", "G#\nAb'", "A'", "A#\nBb'", "B'"];

document.addEventListener('DOMContentLoaded', () => {
    const piano = document.getElementById('piano');
    const octaves = 1;

    for (let octave = 0; octave < octaves; octave++) {
        notes.forEach((note, index) => {
            const key = document.createElement('div');
            key.className = `key ${note.includes('#') ? 'black' : 'white'}`;
            key.textContent = note;
            
            if (!note.includes('#')) {
                piano.appendChild(key);
            } else {
                key.style.left = `${40 * (index + octave * 12) - 15}px`;
                piano.appendChild(key);
            }
            key.addEventListener('mousedown', () => {
                keyPress(key);
            });
        });
    }
});

//piano_interval.html
function keyRelease(key) {
    key.classList.remove('active');
    activeKeyOrder = activeKeyOrder.filter(note => note !== key.dataset.note); // Remove from order list on release
}

function autoReleaseFirstPressedKey() {
    const firstKeyNote = activeKeyOrder.shift(); // Remove the first element from the order array
    const firstKey = keys.find(key => key.dataset.note === firstKeyNote);
    firstKey.classList.remove('active'); // Remove 'active' class from the first pressed key
}

function updateDisplay() {
    const activeKeys = keys.filter(key => key.classList.contains('active'));
    if (activeKeys.length > 0) {
        keyDisplay.textContent = 'Keys pressed: ' + activeKeyOrder.join(', ');
    } else {
        keyDisplay.textContent = 'No keys pressed';
        activeKeyOrder = []; 
    }
}