export const availableLangs = [
  {lang: 'golang', value: 'go', name: 'Go 1.8', default: true},
  {lang: 'python', value: 'python', name: 'Python3', default: false},
  {lang: 'rust', value: 'rust', name: 'Rust', default: false},
  // {lang: 'c', value: 'c', name: 'c', default: false},
  // {lang: 'c++', value: 'c++', name: 'c++', default: false},
  // {lang: 'java', value: 'java', name: 'Java', default: false},
  // {lang: 'javascript', value: 'javascript', name: 'Javascript', default: true},
]

export const langBeMap2Fe = {
  'golang': 'go',
  'python3': 'python',
  'rust': 'rust',
  'c': 'c',
  'c++': 'c++',
  'javascript': 'javascript',
  'java': 'java',
}

export const langFeMap2Be = {
  'go': 'golang',
  'python': 'python3',
  'rust': 'rust',
  'c': 'c',
  'c++': 'c++',
  'javascript': 'javascript',
  'java': 'java',
}