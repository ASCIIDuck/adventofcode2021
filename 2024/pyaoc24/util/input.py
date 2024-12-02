import requests


def transform_input(transformer: callable, args: list = []) -> callable:
    def param_wrapper(decorated_function: callable) -> callable:
        def puzzle_parse(puzzle: str) -> int:
            parsed_args = transformer(puzzle.strip(), *args)
            return decorated_function(*parsed_args)

        return puzzle_parse

    return param_wrapper


def get_token() -> str:
    try:
        with open(".token", "r") as f:
            return f.read().strip()
    except:
        return ""


def get_input(year: int, day: int) -> str:
    cache_path = ".cache/day%02d.input" % day
    try:
        with open(cache_path, "+r") as f:
            return f.read().strip()
    except FileNotFoundError:
        url = "https://adventofcode.com/%d/day/%d/input" % (year, day)
        try:
            sess = requests.Session()
            resp = sess.get(url, cookies={"session": get_token()})
            resp.raise_for_status()
            puzzle_input = resp.content.decode()
            with open(cache_path, "+w") as f:
                f.write(puzzle_input)
        except requests.exceptions.HTTPError:
            print("Failed to fetch input")
            exit(1)
        return puzzle_input
