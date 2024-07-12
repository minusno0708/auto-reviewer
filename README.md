# auto-reviewer

## Usage
- アクセストークンの設定
PRにコメントするために、GITHUBのアクセストークンを設定する

1. アクセストークンを生成する
    ユーザーアイコンをクリックし,Setting/Developer settings/Personal access tokens/Tokens(classic)にアクセスする.

    Generate new tokenクリックし,Expirationから有効期限を設定した後,repo,workflowにチェックを入れてトークンを生成する.

    生成したトークンは,別の場所に控えておく.


2. リポジトリにアクセストークンを設定する

    目的のリポジトリからSettings/Secrets and variables/Actionsにアクセスする.

    New Repository secretをクリックして,NameにAccess_Tokenと入力し,Secretに先ほど生成したトークンをペーストする.

    Add secretから設定を完了する.
