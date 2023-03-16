def is_bool(s: str) -> bool:
    """
    Check if the given string can be cast to boolean

    :param s: string to check
    :return: True if s is a boolean value, False otherwise
    """
    return s.lower() in ["true", "false"]


def str_to_bool(s: str) -> bool:
    """
    Convert the given string to boolean.

    :raise: if given string cannot be cast to boolean.
    :param s: string to convert
    :return: string converted to boolean
    """

    if s in ["true", "True"]:
        return True
    if s in ["false", "False"]:
        return False

    raise Exception(f'{s} is not a string boolean value.')