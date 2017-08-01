from django.shortcuts import render, redirect
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
            user_login(request, username, password)
    return HttpResponseRedirect('/')


def user_logout(request):
    logout(request)
    return HttpResponseRedirect('/')


def user_login(request, username, password):
    user = authenticate(request, username=username, password=password)
    if user is not None:
        login(request, user)


def user_register(request):
    if request.method == 'POST':
        form = RegForm(request.POST)
        if form.is_valid():
            login = request.POST['login']
            email = request.POST['email']
            password = request.POST['password']
            re_password = request.POST['re_password']
            if password == re_password:
                user = User.objects.create_user(login, email, password)
                user.save()
                user_login(request, login, password)
    return redirect('botFactory:index')
