const Message = function() {
  this.text = ''
  this.user = ''
}
Message.prototype.isValid = function () {
  return this.text !== '' && this.user !== ''
}

Vue.component('message', {
  props: ['user', 'text'],
  template: '<div>{{ text }} - {{ user }}</div>'
})

const app = new Vue({
  el: '#app',
  data: {
    messages: [
      {'user': 'nk5', 'text': '1dayインターン始まるよ'},
      {'user': 'saxsir', 'text': '何やるですか'},
      {'user': 'nk5', 'text': 'Go, Vue.js'},
      {'user': 'saxsir', 'text': 'えー'},
    ],
    newMessage: new Message()
  },
  methods: {
    sendMessage() {
      const message = this.newMessage;
      if (!message.isValid()) {
        alert('Input cannot be blank')
        return
      }
      this.messages.push(message)
      this.clearMessage()
    },
    clearMessage() {
      this.newMessage = new Message()
    }        
  }
})