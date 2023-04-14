import openai

# use ChatGPT to translate English to Chinese
def translate(data):
    if data == "":
        return data
    # prompt = "I want you to act as an Chinese translator. I will speak to you in English and you will translate it to Chinese. Do not output anything other than the translation."
    prompt = "将以下文字翻译成中文，翻译得自然、流畅和地道。无法翻译时输出ERR"
    output = openai.ChatCompletion.create(
    model='gpt-3.5-turbo',
    messages=[
            {"role": "system", "content": prompt},
            {"role": "user", "content": data},
        ],
        temperature=0
    )
    return output.choices[0]['message']['content']

if __name__ == '__main__':
    print(translate("Tope"))
    print(translate("When did Virgin Australia start operating?"))
    # print(translate("One of the key advantages of ChatGPT over popular translation tools like Google Translate is its ability to accurately consider the context of a text when generating translations. Considering context can be the difference between simply translating individual words in a sentence and generating a translation that truly reflects the author's or speaker's intention."))
    # print(translate("Hail to thee, blithe Spirit! Bird thou never wert,. That from Heaven, or near it,. Pourest thy full heart. In profuse strains of unpremeditated art."))