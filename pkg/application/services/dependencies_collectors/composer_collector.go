package dependencies_collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

type ComposerCollector struct{}

func NewComposerCollector() *ComposerCollector {
	return &ComposerCollector{}
}

func (c *ComposerCollector) Satisfy(config map[string]interface{}, reference tokens_references.TokenReferenceInterface) (bool, error) {
	if !utils.MapKeyExists(config, "composerPath") || !utils.MapKeyIsString(config, "composerPath") {
		return false, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("ComposerCollector needs the path to the composer.json file_supportive as string.")
	}
	if !utils.MapKeyExists(config, "composerLockPath") || !utils.MapKeyIsString(config, "composerLockPath") {
		return false, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("ComposerCollector needs the path to the composer.lock file_supportive as string.")
	}
	if !utils.MapKeyExists(config, "packages") || !utils.MapKeyIsArrayOfStrings(config, "packages") {
		return false, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("ComposerCollector needs the list of packages as strings.")
	}

	// TODO: implement go.mod parsing
	return false, nil

	//        try {
	//            $this->parser[$config['composerLockPath']] ??= new \Qossmic\Deptrac\Core\Layer\Collector\ComposerFilesParser($config['composerLockPath']);
	//            $parser = $this->parser[$config['composerLockPath']];
	//        } catch (RuntimeException $exception) {
	//            throw new CouldNotParseFileException('Could not parse composer files.', 0, $exception);
	//        }
	//        try {
	//            $namespaces = $parser->autoloadableNamespacesForRequirements($config['packages'], \true);
	//        } catch (RuntimeException $e) {
	//            throw InvalidCollectorDefinitionException::invalidCollectorConfiguration(\sprintf('ComposerCollector has a non-existent package defined. %s', $e->getMessage()));
	//        }
	//        $token = $reference->getToken()->toString();
	//        foreach ($namespaces as $namespace) {
	//            if (\str_starts_with($token, $namespace)) {
	//                return \true;
	//            }
	//        }
	//        return \false;
}
