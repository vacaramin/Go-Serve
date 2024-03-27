document.getElementById('message-form').addEventListener('submit', function(event) {
    event.preventDefault(); 
    
    const messageInput = document.getElementById('message-input');
    const messageText = messageInput.value.trim();

    if (messageText) {
        const newMessage = document.createElement('div');
        newMessage.classList.add('chat-container', 'user');
        newMessage.textContent = messageText;

        const chatBox = document.querySelector('.chat-box');
        chatBox.insertBefore(newMessage, chatBox.lastElementChild);

        messageInput.value = '';

        chatBox.scrollTop = chatBox.scrollHeight;
    }
});