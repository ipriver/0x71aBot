from django import forms


class SignInForm(forms.Form):
    login = forms.CharField(label='  Логин:', min_length=5, max_length=50)
    password = forms.CharField(label='Пароль:', min_length=5, max_length=50)


class RegForm(forms.Form):
    login = forms.CharField(label='Логин: ', min_length=5, max_length=50)
    email = forms.EmailField(label='E-mail:', min_length=3, max_length=50)
    password = forms.CharField(label='Пароль:', min_length=5, max_length=50)
    re_password = forms.CharField(label='Пароль еще раз:', min_length=5,
                                  max_length=50)
