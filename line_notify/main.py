import argparse

from utils.message import send_message


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--message", type=str, help="Message to send")

    # parse input
    args = parser.parse_args()
    message = args.message

    # send message
    r = send_message(message)
    print(r)


if __name__ == "__main__":
    main()
