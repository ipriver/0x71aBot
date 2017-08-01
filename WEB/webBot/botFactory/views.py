from django.shortcuts import render, redirect
from .forms import SignInForm, RegForm, AddBotForm
from django.http import HttpResponseRedirect
from django.contrib.auth import logout
from django.contrib.auth.models import User
from django.contrib.auth import authenticate, login
from .models import Account, Bot
from django.views.decorators.csrf import csrf_exempt


def index(request):
    template_name = 'botFactory/index.html'
    context = {
        'authenticated': False,
        'form': None,
    }
    if request.user.is_authenticated:
        context['message'] = 'Hello, ' + request.user.username
        context['authenticated'] = True
        context['form'] = AddBotForm()
        user = request.user
        account = Account.objects.get(user=user)
        bot_list = Bot.objects.filter(account=account)
        context['bot_list'] = bot_list
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
                account = Account(user=user, bot_count=0)
                account.save()
                user_login(request, login, password)
    return redirect('botFactory:index')


def add_new_bot(request):
    if request.method == 'POST':
        form = AddBotForm(request.POST)
        if form.is_valid():
            channel_name = request.POST['channel_name']
            user = request.user
            account = Account.objects.get(user=user)
            if account.bot_count < 5:
                bot = Bot(account=account, channel_name=channel_name)
                bot.bot_start()
                account.bot_count += 1
                account.save()
                bot.save()
    return redirect('botFactory:index')


@csrf_exempt
def del_bot(request):
    if request.is_ajax:
        user = request.user
        account = Account.objects.get(user=user)
        account.bot_count -= 1
        bot_list = Bot.objects.filter(account=account)
        bot = Bot.objects.all()
        bot.filter(id=bot_list[0].id).delete()
        account.save()

    return redirect('botFactory:index')
