import requests


def transformInput(transformer: callable, args: list = []) -> callable:
    def paramWrapper(decorated_function: callable) -> callable:
        def puzzleParse(puzzle: str) -> int:
            parsed_args = transformer(puzzle, *args)
            return decorated_function(*parsed_args)

        return puzzleParse

    return paramWrapper


def getToken() -> str:
    try:
        with open(".token", "r") as f:
            return f.read().strip()
    except:
        return ""


def getInput(year: int, day: int) -> str:
    cache_path = ".cache/day%02d.input" % day
    try:
        with open(cache_path, "+r") as f:
            return f.read().strip()
    except FileNotFoundError:
        url = "https://adventofcode.com/%d/day/%d/input" % (year, day)
        sess = requests.Session()
        resp = sess.get(url, cookies={"session": getToken()})
        input = resp.content.decode()
        with open(cache_path, "+w") as f:
            f.write(input)
        return input
