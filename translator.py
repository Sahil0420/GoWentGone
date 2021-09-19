import googletrans
import speech_recognition as sr
import gtts
import playsound
from tkinter import*
import os
def show():
	label.config(text="Selected Language : "+clicked.get())

def clear(event):
	scvalue.set("")
	scvalue1.set("")	


def translate():
	global tt
	global text
	reco=sr.Recognizer()
	try:
		with sr.Microphone() as source:
			print('Speak now')
			voice=reco.listen(source)
			text=reco.recognize_google(voice)
			print(text)
			scvalue1.set("Text :"+text)
			res1.update()
			her_lang=clicked.get()
			trans=googletrans.Translator()
			translated=trans.translate(text,dest=her_lang)
			tt=(translated.text)
			print(tt)
			scvalue.set("Translated text :"+tt)
			res.update()
			converted_audio=gtts.gTTS(translated.text,lang=her_lang)
			converted_audio.save('test_voice.mp3')
			playsound.playsound("test_voice.mp3")
	except:
		pass
#DROP DOWN MENUS:
options=[
"en","es","fr","ko","jp","hi"]
win=Tk()
win.configure(bg="white")
 
clicked=StringVar()
clicked.set("Select Language")
scvalue=StringVar()
scvalue.set("")
scvalue1=StringVar()
scvalue1.set("")

win.geometry("300x500")
f=Frame(win,height="250",width="200",bg='white')
header=Label(f,text="TRANSLATOR",font=("Helvecita",20,"bold"),bg='white',fg="black")
header.pack()
btn=Button(f,text="Speak",width="8",height="3",bg="aqua",fg="black",command=translate,font=("lucida",10,"bold"),borderwidth=4)
btn.pack(side=LEFT,padx=5,pady=5)
btn=Button(f,text="Quit",width="8",height="3",bg="red",fg="#ebeae8",font=("lucida",10,"bold"),borderwidth=4)
btn.bind("<Button-1>",quit)
btn.pack(side=RIGHT,padx=5,pady=5)
f.pack(padx=5,pady=5)

f2=Frame(win,height="50",width="200",bg='white')

lbx=Listbox(f2,fg='blue',bg='white')
lbx.pack()
lbx.insert(ACTIVE,"...LANGUAGE LIST....")
lbx.insert(END,"1. en=English")
lbx.insert(END,"2. es=Spanish")
lbx.insert(END,"3. fr=French")
lbx.insert(END,"4. Ko=korean")
lbx.insert(END,"5. ja=Japanese")
lbx.insert(END,"6. hi=Hindi")
label=Label(f2,text=" ")
label.pack()
f2.pack(padx=10,pady=5)

f3=Frame(win,height="250",width="200",bg='white',)
res1=Entry(f3,textvar=scvalue1,bg="aqua",font=('helvecita',10),width="100",borderwidth=4)
res1.pack(pady=10)

res=Entry(f3,textvar=scvalue,bg='aqua',font=('lucida',10),width="100",borderwidth=4)
res.pack(pady=10)

drop=OptionMenu(f2,clicked,*options)
drop.pack(padx=5,pady=10)



btn=Button(f3,text="Clear",width="10",height="3",font=("lucida",10,"bold"),bg="red",borderwidth=5)
btn.pack()
btn.bind("<Button-1>",clear)
f3.pack()
#print(googletrans.LANGUAGES)
win.mainloop()