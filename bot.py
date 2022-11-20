import telebot
import json
import subprocess


bot = telebot.TeleBot('5700965158:AAF254qjyjXlAmuJGbe51ce8jKz_JM1VD_s')

global history
history = 0

@bot.message_handler(content_types=['text'])
def get_text_messages(message):
    if message.text == "/start":
            bot.send_message(message.from_user.id, "Hi!\nI will help you find reliable sources of information and show you which one has fakes.\nTo start your journey, write /get")
    elif message.text == "/get":
        history += 1
        command = './hackaton-tool --url=https://krakowexpats.pl/tips-articles/page/{0}'.format(history)
        subprocess.run([command],shell=True)
        
        with open('db.json', 'r') as database:
            database = json.load(database)
        
        i = 0
        while i < len(database):
            source = database[i]["Source"] + "\n"
            author = database[i]["Author"] + "\n"
            title = database[i]["Title"]
            href = database[i]["Href"] + "\n"
            publication_time = database[i]["Publication_time"] +"\n"
            category = database[i]["Category"] +"\n"
            fake_detector = database[i]["Fake_detector"] +"\n"
            mood = database[i]["Mood"] +"\n"
            i += 1
            
            text = '«{0}»\n\nAuthor: {1}Publication time: {2}Source: {3}Link: {4}Category: {5}Mood: {6}Checking for fake: {7}'.format(
                title, author, publication_time, source, href, category, mood, fake_detector
            )
            
            bot.send_message(message.from_user.id, text)
    elif message.text == "/help":
        bot.send_message(message.from_user.id, "Coming soon...")
    else:
        pass

bot.polling(none_stop=True, interval=0)