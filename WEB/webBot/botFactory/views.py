from django.shortcuts import render
from .forms import SignInForm, RegForm
from django.http import HttpResponseRedirect
from django.contrib.auth import logout
from django.contrib.auth.models import User
from django.contrib.auth import authenticate, login


def index(request):
    template_name = 'botFactory/index.html'
    context = {
        'authenticated': False,
        'form': None,
    }
    if request.user.is_authenticated:
        context['message'] = 'Hello, ' + request.user.username
        context['authenticated'] = True
    else:
        context['form'] = SignInForm()
    return render(request, template_name, context)


def register(request):
    template_name = 'botFactory/register.html'
    context = {
        'form': RegForm()
    }
    return render(request, template_name, context)


def sign_in(request):
    if request.method == 'POST':
        form = SignInForm(request.POST)
        if form.is_valid():
            username = request.POST['login']
            password = request.POST['password']
            user = authenticate(request, username=username, password=password)
            if user is not None:
                login(request, user)
            return HttpResponseRedirect('/')
    else:
        form = SignInForm()
        context = {
            'authenticated': False,
            'form': form,
        }
    return render(request, template_name, context)

def user_logout(request):
    logout(request)
    return HttpResponseRedirect('/')