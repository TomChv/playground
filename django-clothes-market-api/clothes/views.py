import json

from django.http import HttpResponse, HttpRequest, JsonResponse
from django.shortcuts import get_object_or_404
from rest_framework.viewsets import ViewSet

from common.exceptions import BadRequest
from common.utils import is_bool, str_to_bool

from .models import Article, ArticleSerializer, CreateArticleValidator, UpdateArticleValidator


def ping(_: HttpRequest) -> HttpResponse:
    return HttpResponse(status=200)


class ArticleView(ViewSet):
    def list(self, request: HttpRequest) -> BadRequest | JsonResponse:
        # Check if sold_filter is set and its type is correct
        sold_filter = request.GET.get(key="sold", default=None)

        if sold_filter is not None:
            if is_bool(sold_filter) is False:
                return BadRequest(message="invalid sold filter type, it shall be a boolean.").res()
            articles = list(Article.objects.filter(sold=str_to_bool(sold_filter)).values())
        else:
            articles = list(Article.objects.values())

        return JsonResponse({
            "number_of_articles": len(articles),
            "articles": articles
        }, safe=False)

    def retrieve(self, request, pk=None):
        articles = Article.objects.all()
        article = get_object_or_404(articles, pk=pk)

        serializer = ArticleSerializer(article)
        return JsonResponse(serializer.data)

    def create(self, request: HttpRequest):
        body_unicode = request.body.decode('utf-8')
        body = json.loads(body_unicode)

        parsed_body = CreateArticleValidator(data=body)
        if parsed_body.is_valid() is False:
            return BadRequest(message=f'error in the body: {parsed_body.error_messages}').res()

        new_article = Article(
            name=parsed_body.data['name'],
            state=parsed_body.data['state'],
            size=parsed_body.data['size'],
            brand=parsed_body.data['brand'],
            image_url=parsed_body.data['image_url'],
            description=parsed_body.data['description'],
            buy_price=parsed_body.data['buy_price'])

        new_article.save()

        return JsonResponse(ArticleSerializer(new_article).data)

    def update(self, request: HttpRequest, pk=None):
        articles = Article.objects.all()
        article = get_object_or_404(articles, pk=pk)

        body_unicode = request.body.decode('utf-8')
        body = json.loads(body_unicode)

        parsed_body = UpdateArticleValidator(data=body)
        if parsed_body.is_valid() is False:
            return BadRequest(message=f'error in the body: {parsed_body.error_messages}').res()

        # Check if the property shall be updated and set it
        # There is not safer way to do it than a list of if
        # statement (looping through map is not possible
        # because article properties cannot be accessed as
        # dictionary key
        if "name" in parsed_body.data:
            article.name = parsed_body.data['name']

        if "image_url" in parsed_body.data:
            article.image_url = parsed_body.data['image_url']

        if "size" in parsed_body.data:
            article.size = parsed_body.data['size']

        if "state" in parsed_body.data:
            article.state = parsed_body.data['state']

        if "sold" in parsed_body.data:
            article.sold = parsed_body.data['sold']

        if "sold_price" in parsed_body.data:
            article.sold_price = parsed_body.data['sold_price']

        if "brand" in parsed_body.data:
            article.brand = parsed_body.data['brand']

        if "description" in parsed_body.data:
            article.description = parsed_body.data['description']

        if "buy_price" in parsed_body.data:
            article.buy_price = parsed_body.data['buy_price']

        article.save()

        return JsonResponse(ArticleSerializer(article).data)

    def destroy(self, request: HttpRequest, pk=None):
        articles = Article.objects.all()
        article = get_object_or_404(articles, pk=pk)

        article.delete()

        return JsonResponse(ArticleSerializer(article).data)
