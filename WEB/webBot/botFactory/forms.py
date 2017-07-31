from django import forms


class SignInForm(forms.Form):
    login = forms.CharField(label='Логин', min_length=5, max_length=50)
    password = forms.CharField(label='Пароль:', min_length=5, max_length=50)
