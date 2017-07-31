from django.shortcuts import render
from .forms import SignInForm
#from django.http import HttpResponse


def test(request):
    template_name = 'botFactory/index.html'
    context = {
        'message': '',
        'authenticated': False,
        'form': None,
    }
    if request.user.is_authenticated:
        context['message'] = 'Hello, user'
        context['authenticated'] = True
    else:
        context['message'] = 'Анонимный пользователь'
        context['form'] = SignInForm()
    return render(request, template_name, context)


def register(request):
    pass
